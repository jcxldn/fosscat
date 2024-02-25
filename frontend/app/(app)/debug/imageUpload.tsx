import React from "react";
import { View } from "../../../components/Themed";
import { Button, Text } from "react-native-paper";

import * as ImagePicker from 'expo-image-picker';
import { fetchImageFromUri } from "../../../util/fetchImageFromUri";
import { useSession } from "../../../components/AuthenticationContext";

import ReactNativeBlobUtil from 'react-native-blob-util'

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
                        console.log(result.assets[i].uri)
                        const blob = await fetchImageFromUri(result.assets[i].uri)
                        console.log(blob.size)
                        formData.append("upload[]", blob)
                    }
                }

                console.log(formData)

                const res = await ReactNativeBlobUtil.fetch(
                    "POST",
                    "http://10.255.0.202:8080/upload",
                    {
                        "Authorization": `Bearer ${jwt}`,
                        // this is required, otherwise it won't be process as a multipart/form-data request
                        'Content-Type': 'multipart/form-data',
                    },
                    [
                        result.assets?.map(asset => ({
                            name: "upload[]",
                            filename: asset.fileName,
                            data: ReactNativeBlobUtil.wrap(asset.uri)
                        }))
                    ],
                ).catch((err: Error) => {
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
        </View >
    )
}

export default ImageUploadPage;