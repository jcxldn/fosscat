import React from "react";

import { Redirect, Stack } from "expo-router";
import { useSession } from "../../components/AuthenticationContext";
import { Text } from "../../components/Themed";
import DrawerComponent from "../../components/Drawer";

const AppLayout = () => {
    const { getCreds } = useSession()
    const [creds, setCreds] = React.useState(-1)


    // React effects cannot be async, have self calling async block inline instead.
    React.useEffect(() => {
        const fn = (async () => {
            const hasCreds = await getCreds()
            // Cast to number so we can have starting value (-1) and boolean values (0, 1)
            setCreds(hasCreds as unknown as number);

        });
        if (creds == -1) fn();

    })

    if (creds == -1) {
        return <Text>Loading...</Text>
    } else if (creds == 1) { // creds == true
        return <DrawerComponent />
    } else {
        return <Redirect href="/login" />
    }
}

export default AppLayout;