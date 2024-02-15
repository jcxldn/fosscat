// Based on https://github.com/expo/examples/blob/master/with-aws-storage-upload/App.js#L103

export const fetchImageFromUri = async (uri: string) => {
    const response = await fetch(uri);
    const blob = await response.blob();
    return blob;
};