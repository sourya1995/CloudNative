module.exports = class RequestProcessor {
    constructor(uploadSigner, downloadSigner, uploadLimitInMB, allowedExtensions) {
        this.uploadSigner = uploadSigner;
        this.downloadSigner = downloadSigner;
        this.uploadLimitInMB = uploadLimitInMB;
        this.allowedExtensions = allowedExtensions;
};

processRequest(requestId, extension) {
    if(!extension) {
        throw 'no extension specified';
    }
    const normalizedExtension = extension.toLowerCase();
    const isImage = this.allowedExtensions.includes(normalizedExtension);
    if(!isImage) {
        throw 'extension ${extension} not supported';
    }
    const fileKey = `${requestId}.${normalizedExtension}`;
    return {
        upload: this.uploadSigner.signUpload(fileKey, this.uploadLimitInMB),
        download: this.downloadSigner.signDownload(fileKey)
    };
};
};