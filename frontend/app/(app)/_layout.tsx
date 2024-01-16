import React from "react";

import { Redirect, Stack } from "expo-router";
import { useSession } from "../../components/AuthenticationContext";
import { Text } from "../../components/Themed";

const AppLayout = () => {
    const { jwt, isLoading } = useSession()

    if (isLoading) {
        return <Text>Loading...</Text>
    }

    if (!jwt) {
        return <Redirect href="/login" />
    }

    return <Stack />
}

export default AppLayout;