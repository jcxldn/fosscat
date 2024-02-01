import React from "react";
import { Text } from "../../components/Themed";
import { Stack } from "expo-router";

const HomePage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "Home" }} />
            <Text>Hello, world!</Text>
        </>
    )
}

export default HomePage;