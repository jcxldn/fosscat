import React from "react";
import { View } from "../../../components/Themed";
import { Button, Text } from "react-native-paper";

import * as ImagePicker from 'expo-image-picker';
import { useSession } from "../../../components/AuthenticationContext";
import { useApolloClient } from "../../../components/ApolloClientProvider";

const ImageUploadPage = () => {
    const jwt = useSession().jwt
    const { serverUri } = useApolloClient();

    return (
        <View>
            <Button onPress={async () => {
                const result = await ImagePicker.launchImageLibraryAsync({
                    mediaTypes: ImagePicker.MediaTypeOptions.Images,
                    allowsEditing: true,
                    quality: 1
                })
                const formData = new FormData();

                // If assets exists, iterate over it
                if (result.assets) {
                    for (let i = 0; i < result.assets?.length; i++) {
                        const asset = result.assets[i]
                        // Append to formData an object contaning the file uri, mimetype and name
                        // These are the required values for it to be detected as a file.
                        // on iOS, we must use this method (as opposed to uploading a blob using
                        // `formData.append("xyz", blob)`) as otherwise the file contents are not present in the request
                        formData.append("upload[]", {
                            uri: asset.uri, // usually "file://" in this context.
                            type: asset.mimeType || "application/octet-stream",
                            // Filename is required or the form entry will not be recognised as a file
                            name: asset.fileName || asset.uri.split("/").pop()
                            // Cast the object to a Blob to make TypeScript happy
                            // See: https://github.com/g6ling/React-Native-Tips/issues/1#issuecomment-1165945160
                        } as unknown as Blob)
                    }
                }

                const res = await fetch(`${serverUri}/upload`, {
                    method: "POST",
                    headers: {
                        "Authorization": `Bearer ${jwt}`
                    },
                    body: formData,
                }).catch((err: Error) => {
                    // TODO, make cleaner, just a proof of concept.
                    alert("Error uploading image: " + err.message)
                })

                // Check if res is of type Response
                // If not, .catch would have been called so no additional handling is required.
                if (res instanceof Response) {
                    if (res.status == 200) {
                        // Uploaded successfully
                        alert("Uploaded sucessfully")
                    } else {
                        // Did not upload successfully.
                        alert("Error uploading image: Response was not 200")
                        console.log(await res.text())
                        console.log(formData)
                    }
                }
            }}>
                <Text>Upload File</Text>
            </Button>
        </View>
    )
}

export default ImageUploadPage;