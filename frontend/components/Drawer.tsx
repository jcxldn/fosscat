import * as React from 'react';

import { Drawer } from 'expo-router/drawer';
import { DrawerContentScrollView, DrawerHeaderProps, DrawerItem, useDrawerStatus } from '@react-navigation/drawer';
import { View, StyleSheet } from 'react-native';
import { Appbar, Avatar, Caption, Title } from 'react-native-paper';
import { router } from 'expo-router';


// layout prop not used
const Header = ({ navigation, options, route }: DrawerHeaderProps) => {

    const isDrawerOpen = useDrawerStatus()
    return (
        <Appbar.Header>
            {
                // If can go back (and menu is not open)
                // If we only check if we can go back it switches from menu to back icon when the drawer is opened.
                navigation.canGoBack() && !isDrawerOpen ?
                    // Can go back, show back button
                    <Appbar.BackAction onPress={navigation.goBack} /> :
                    // Can't go back, show menu button for drawer
                    <Appbar.Action icon="menu" onPress={navigation.openDrawer} />
            }

            {/** Set header title to options.title if present, else route name */}
            <Appbar.Content title={options.title ? options.title : route.name} />
        </Appbar.Header>
    )
}

const DrawerContent = () => {
    return (
        <DrawerContentScrollView>
            {/** Header section showing user info */}
            <View style={styles.drawerHeaderSection}>
                <View style={styles.horizontalAlign}>
                    <Avatar.Image source={require("../assets/images/Icon_1024x1024.png")} />
                    <View style={{ paddingLeft: 16 }}>
                        <Title>James Cahill</Title>
                        <Caption>example@example.com</Caption>
                    </View>
                </View>
            </View>

            {/** Route buttons */}
            <DrawerItem label="Home" onPress={() => { router.navigate("/") }} />
            <DrawerItem label="Scan" onPress={() => { router.navigate("/scan") }} />
            <DrawerItem label="About" onPress={() => { router.navigate("/about") }} />
        </DrawerContentScrollView>
    )
}

// TODO: https://docs.expo.dev/router/advanced/platform-specific-modules/
const DrawerComponent = () => {
    return (
        <Drawer
            id="root"
            drawerContent={DrawerContent}
            screenOptions={{ header: Header }}
        />
    );
};

const styles = StyleSheet.create({
    drawerHeaderSection: {
        paddingTop: 16,
        paddingLeft: 16,
        paddingBottom: 16
    },
    horizontalAlign: {

        display: "flex",
        alignItems: "center",
        // justifyContent: "center",
        flexDirection: "row",
    }
})

export default DrawerComponent;