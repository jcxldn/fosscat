import React from "react";

import { Redirect, Stack } from "expo-router";
import { useSession } from "../../components/AuthenticationContext";
import { Text } from "../../components/Themed";
import DrawerComponent from "../../components/Drawer";

const AppLayout = () => {
    const { jwt, isLoading } = useSession()

    if (isLoading) {
        return <Text>Loading...</Text>
    }

    if (!jwt) {
        return <Redirect href="/login" />
    }

    return (
        <DrawerComponent />
    )
}

export default AppLayout;