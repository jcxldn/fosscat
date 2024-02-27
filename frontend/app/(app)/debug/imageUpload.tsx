import React from "react";
import { View } from "../../../components/Themed";
import { Button, Text } from "react-native-paper";

import * as ImagePicker from 'expo-image-picker';
import { fetchImageFromUri } from "../../../util/fetchImageFromUri";
import { useSession } from "../../../components/AuthenticationContext";

const ImageUploadPage = () => {
    const jwt = useSession().jwt

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
                        console.log("Adding asset")
                        const asset = result.assets[i]
                        formData.append("upload[]", {
                            uri: asset.uri,
                            type: asset.mimeType || "application/octet-stream",
                            // Filename is required or the form entry will not be recognised as a file
                            name: asset.fileName || asset.uri.split("/").pop()
                            // https://github.com/g6ling/React-Native-Tips/issues/1#issuecomment-1165945160
                        } as unknown as Blob)
                    }
                }

                console.log(formData)

                const res = await fetch("http://172.20.10.8:8080/upload", {
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