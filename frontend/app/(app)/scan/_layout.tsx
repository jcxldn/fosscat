import React from "react";
import { Button } from "react-native";
import { Stack, router } from "expo-router";

export default function ScanStack() {
    return (
        <Stack id="stack.scan">

            <Stack.Screen name="camera" options={{
                headerShown: false
            }} />

            <Stack.Screen name="result" options={{
                headerTitle: "Scan Result",
                headerBackButtonMenuEnabled: true,

                headerLeft: ({ label }) => (
                    <Button onPress={() => router.back()} title="Close" />
                ),
                presentation: "modal"

            }} />
        </Stack>
    )
}