import React from "react";
import { View, Text } from "../components/Themed";
import { Stack } from "expo-router";

import { TextInput } from 'react-native-paper';
import { StyleSheet } from "react-native";

const LoginPage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "Login" }} />
            <View style={styles.container}>
                <Text style={styles.title}>Login</Text>
                <View style={styles.itemContainer}>
                    <TextInput label="Email" />
                </View>
                <View style={styles.itemContainer}>
                    <TextInput label="Password" />
                </View>
            </View>
        </>
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
        paddingTop: 5,
        paddingBottom: 5
    }
})

export default LoginPage;