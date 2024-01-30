import React from "react";
import Constants from "expo-constants";
import { ApplicationReleaseType, applicationId, applicationName, getIosApplicationReleaseTypeAsync, nativeApplicationVersion, nativeBuildVersion } from "expo-application";
import { Stack, router } from "expo-router";

import { Platform, StyleSheet } from "react-native";
import { Avatar, Button, Divider } from "react-native-paper";


import { Text } from 'react-native-paper';

import { View } from "../../components/Themed";

const IosAppReleaseType = () => {
    const [releaseType, setReleaseType] = React.useState<ApplicationReleaseType | null>(null);

    React.useEffect(() => {
        const func = async () => {
            setReleaseType(await getIosApplicationReleaseTypeAsync());
        };
        func()
    });
    if (!releaseType) {
        return <Text>Fetching release type...</Text>
    }

    let releaseTypeStr = ""
    switch (releaseType) {
        case ApplicationReleaseType.SIMULATOR:
            releaseTypeStr = "Simulator"
            break;
        case ApplicationReleaseType.ENTERPRISE:
            releaseTypeStr = "Enterprise"
            break;
        case ApplicationReleaseType.DEVELOPMENT:
            releaseTypeStr = "Development"
            break;
        case ApplicationReleaseType.AD_HOC:
            releaseTypeStr = "Ad-hoc"
            break;
        case ApplicationReleaseType.APP_STORE:
            releaseTypeStr = "App Store"
            break;
        default:
            releaseTypeStr = "Unknown"
            break;
    }

    return (
        <Text>{releaseTypeStr} Build</Text>
    )

}

const AboutPage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "About" }} />
            <View style={{ flex: 1, padding: 16 }}>
                <View style={styles.banner}>
                    {/** Default size 64 */}
                    <Avatar.Image source={require("../../assets/images/Icon_1024x1024.png")} />
                    <View style={styles.bannerText}>
                        <Text variant="displaySmall">{applicationName}</Text>
                        <Text variant="labelLarge">{applicationId}</Text>
                    </View>
                </View>

                <Divider />
                <Text style={styles.headline} variant="headlineSmall">Developer</Text>
                {Platform.OS == "ios" ? <IosAppReleaseType /> : undefined}
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

const styles = StyleSheet.create({
    banner: {
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        flexDirection: "row",
        paddingTop: 16, /** already have 16 padding on page itself */
        paddingBottom: 32
    },
    bannerText: {
        paddingLeft: 32
    },
    headline: {
        paddingTop: 16,
        paddingBottom: 16
    }
})

export default AboutPage;