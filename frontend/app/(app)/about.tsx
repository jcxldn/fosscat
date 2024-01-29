import React from "react";
import Constants from "expo-constants";
import { ApplicationReleaseType, applicationId, applicationName, nativeApplicationVersion, nativeBuildVersion } from "expo-application";
import { Stack, router } from "expo-router";

import { Platform } from "react-native";
import { Button } from "react-native-paper";

import { View, Text } from "../../components/Themed";

const AboutPage = () => {
    let releaseType = "";
    switch (ApplicationReleaseType) {
        case ApplicationReleaseType.SIMULATOR:
            releaseType = "Simulator"
            break;
        case ApplicationReleaseType.ENTERPRISE:
            releaseType = "Enterprise"
            break;
        case ApplicationReleaseType.DEVELOPMENT:
            releaseType = "Development"
            break;
        case ApplicationReleaseType.AD_HOC:
            releaseType = "Ad-hoc"
            break;
        case ApplicationReleaseType.APP_STORE:
            releaseType = "App Store"
        default:
            releaseType = "Unknown"

    }
    return (
        <>
            <Stack.Screen options={{ title: "About" }} />
            <View style={{ flex: 1, padding: 16 }}>
                <Text>{applicationName} ({releaseType}) - {applicationId}</Text>

                <Text>Version {nativeApplicationVersion}</Text>
                <Text>Build {nativeBuildVersion}</Text>
                <Text>Expo SDK {Constants.expoConfig?.sdkVersion}</Text>
                <Text>React {React.version}</Text>
                <Text>React Native {Platform.constants.reactNativeVersion.major}.{Platform.constants.reactNativeVersion.minor}.{Platform.constants.reactNativeVersion.patch}</Text>
                <View style={{ paddingTop: 16 }}>
                    <Text>{Platform.OS} ({Platform.Version})</Text>
                </View>
                <View style={{ paddingTop: 16 }}>
                    <Button onPress={() => {
                        if (Platform.OS == "web") {
                            localStorage.clear()
                        } else {
                            alert("Not on web, could not clear localStorage.")
                        }
                    }}>
                        <Text>Reset localStorage (web)</Text>
                    </Button>
                    <Button onPress={() => router.navigate("/_sitemap")}>
                        <Text>Sitemap</Text>
                    </Button>
                </View>

            </View>
        </>
    )
}

export default AboutPage;