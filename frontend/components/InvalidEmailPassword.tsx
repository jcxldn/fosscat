import React from "react";

import { Text, View } from "./Themed";
import { Surface, useTheme } from "react-native-paper";
import { StyleSheet } from "react-native";

import { MaterialIcons } from '@expo/vector-icons';

export const InvalidEmailPassword = () => {
    const { colors } = useTheme();
    return (
        <View style={style.outer}>
            <Surface style={{
                backgroundColor: colors.errorContainer,
                ...style.surface
            }}>
                <MaterialIcons name="error-outline" size={24} color="black" style={{ paddingRight: 8 }} />
                <Text>Invalid email / password.</Text>

            </Surface>
        </View>
    )
}

const style = StyleSheet.create({
    outer: {
        paddingTop: 16,
        paddingBottom: 16
    },
    surface: {
        padding: 8,
        width: 250,
        alignItems: "center",
        justifyContent: "center",
        flexDirection: "row"
    }
})