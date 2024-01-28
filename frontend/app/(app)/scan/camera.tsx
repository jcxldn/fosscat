import React from "react";
import { View, Text } from "../../../components/Themed";
import { Button, StyleSheet } from "react-native";
import { useNavigation } from "expo-router";
import { BarcodeScanningResult, Camera, CameraView, useCameraPermissions } from "expo-camera/next";

const ScanPage = () => {
    // States
    const [permission, requestPermission] = useCameraPermissions();
    const [scanned, setScanned] = React.useState(false);

    // Hook to grab the parentt navigator (in this case navigator "stack.scan")
    const nav = useNavigation();

    // (fork): https://github.com/jcxldn/expo-camera-barcode-types-error/blob/main/app/modal.tsx
    if (!permission) {
        // Camera permissions are still loading
        return <View />;
    }

    // (fork): https://github.com/jcxldn/expo-camera-barcode-types-error/blob/main/app/modal.tsx
    if (!permission.granted) {
        // Camera permissions are not granted yet
        return (
            <View style={styles.container}>
                <Text style={{ textAlign: "center" }}>
                    We need your permission to show the camera
                </Text>
                <Button onPress={requestPermission} title="Request Permission" />
            </View>
        );
    }

    const handleScan = ({ type, data }: BarcodeScanningResult) => {
        setScanned(true)
        alert(`Barcode scanned with type ${type} and data ${data}.`)
        // Navigate to ./result
        alert(nav.getParent()?.navigate("result", { hello: "here" }))
    }

    return (
        <View style={StyleSheet.absoluteFillObject}>
            <Text>hi</Text>
            <CameraView style={StyleSheet.absoluteFillObject} onBarcodeScanned={scanned ? undefined : handleScan} barcodeScannerSettings={{ barCodeTypes: ["org.iso.Code128", "org.iso.QRCode"] }} />
            {scanned && (
                <Button title="Tap to Scan Again" onPress={() => setScanned(false)} />
            )}
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        height: "100%",
        width: "100%"
    }
});

export default ScanPage;