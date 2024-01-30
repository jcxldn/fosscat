import * as React from 'react';

import { Drawer } from 'expo-router/drawer';


// TODO: https://docs.expo.dev/router/advanced/platform-specific-modules/
const DrawerComponent = () => {
    return (
        <Drawer id="root">
            <Drawer.Screen
                name="index" // url
                options={{
                    drawerLabel: "Home",
                    title: "Home"
                }}
            />
            <Drawer.Screen
                name="scan" // url
                options={{
                    drawerLabel: "Quick Scan",
                    title: "Scan"
                }}
            />
            <Drawer.Screen
                name="about" // url
                options={{
                    drawerLabel: "About",
                    title: "About"
                }}
            />
        </Drawer>
    );
};

export default DrawerComponent;