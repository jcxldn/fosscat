import React from "react";
import { View, Text } from "../components/Themed";
import { router } from "expo-router";

import { Button, TextInput } from 'react-native-paper';
import { KeyboardAvoidingView, Platform, StyleSheet } from "react-native";
import { useSession } from "../components/AuthenticationContext";
import { InvalidEmailPassword } from "../components/InvalidEmailPassword";
import { useApolloClient } from "../components/ApolloClientProvider";

const LoginPage = () => {

    const [email, setEmail] = React.useState("");

    const [password, setPassword] = React.useState("");

    const [showEmailPassError, setShowEmailPassError] = React.useState(false)


    const { login } = useSession();

    const { setServerUri, serverUri } = useApolloClient();

    return (
        <KeyboardAvoidingView style={styles.container} behavior={Platform.OS === "ios" ? "padding" : "height"}>
            <Text style={styles.title}>Login</Text>
            {showEmailPassError ? <InvalidEmailPassword /> : <></>}
            <View style={styles.itemContainer}>
                <View style={styles.itemContainer}>
                    <TextInput
                        label="Server URL"
                        value={serverUri}
                        onChangeText={text => setServerUri(text)}
                    />
                </View>
                <TextInput
                    label="Email"
                    value={email}
                    onChangeText={text => setEmail(text)}
                />
            </View>
            <View style={styles.itemContainer}>
                <TextInput
                    label="Password"
                    value={password}
                    onChangeText={text => setPassword(text)}
                />
            </View>
            <Button
                onPress={async () => {
                    const res = await login(email, password)

                    if (!res?.login.success) {
                        // Login was not sucessful
                        setShowEmailPassError(true)
                    } else {
                        // Login was a success, redirect to homepage.
                        router.replace('/')
                    }
                }}>
                Login
            </Button>
        </KeyboardAvoidingView>
    )
}

const styles = StyleSheet.create({
    title: {
        fontWeight: "bold",
        fontSize: 32,
        paddingBottom: 32
    },
    container: {
        flex: 1,
        alignItems: "center",
        justifyContent: "center",
        padding: 20
    },
    itemContainer: {
        width: "100%",
        paddingBottom: 10
    }
})

export default LoginPage;