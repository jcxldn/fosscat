import React from "react";
import { View, useColorScheme } from "react-native";
import { Button, Text } from "react-native-paper";
import { MaterialCommunityIcons } from '@expo/vector-icons';

const ItemCheckInOutButtons = () => {
    const colorScheme = useColorScheme()
    const buttonColor = colorScheme === "dark" ? "white" : "black"

    return (
        <View style={{ paddingTop: 16, flexDirection: "row", width: "100%", justifyContent: "space-evenly" }}>
            {/** Check in button */}
            <Button mode="elevated" onPress={() => alert("check in tbd")}>
                <View style={{ justifyContent: "center", alignItems: "center" }}>
                    <MaterialCommunityIcons name="clock-in" size={32} color={buttonColor} />
                    <Text>Check in</Text>
                </View>
            </Button>

            {/** Check out button */}
            <Button mode="elevated" onPress={() => alert("check out tbd")}>
                <View style={{ justifyContent: "center", alignItems: "center" }}>
                    <MaterialCommunityIcons name="clock-out" size={32} color={buttonColor} />
                    <Text>Check out</Text>
                </View>
            </Button>
        </View>
    )
}
export default ItemCheckInOutButtons;