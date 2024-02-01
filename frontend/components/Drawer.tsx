import * as React from 'react';

import { Drawer } from 'expo-router/drawer';
import { DrawerContentScrollView, DrawerItem } from '@react-navigation/drawer';
import { View, StyleSheet } from 'react-native';
import { Avatar, Caption, Title } from 'react-native-paper';
import { router } from 'expo-router';


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
        <Drawer id="root" drawerContent={() => <DrawerContent />} />
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