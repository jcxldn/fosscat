import React from "react";
import { Text, View } from "../../components/Themed";
import { BarCodeScanner } from "expo-barcode-scanner";
import { Button, StyleSheet } from "react-native";

const ScanPage = () => {
    // States
    const [hasPermission, setHasPermission] = React.useState<Boolean | null>(null);
    const [scanned, setScanned] = React.useState(false);

    React.useEffect(() => {
        console.log("DEFINE GBSP")
        const getBarcodeScannerPermissions = async () => {
            const { status } = await BarCodeScanner.requestPermissionsAsync();
            setHasPermission(status === "granted")
            console.log(`Status: ${status}`)
        };

        getBarcodeScannerPermissions();
    });

    const handleScan = ({ type, data }) => {
        setScanned(true)
        alert(`Barcode scanned with type ${type} and data ${data}.`)
    }

    if (hasPermission === null) {
        return <Text>Requesting for camera permission</Text>;
    }

    if (!hasPermission) {
        return <Text>Camera permission denied.</Text>
    }

    return (
        <View style={StyleSheet.absoluteFillObject}>
            <Text>hi</Text>
            <BarCodeScanner style={StyleSheet.absoluteFillObject} onBarCodeScanned={scanned ? undefined : handleScan}>
                {scanned && <Button title={'Tap to Scan Again'} onPress={() => setScanned(false)} />}
            </BarCodeScanner>
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