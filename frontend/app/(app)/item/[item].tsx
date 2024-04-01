// Based on https://github.com/callstack/react-native-pager-view/blob/master/example/src/PaginationDotsExample.tsx

import React from "react";
import { View } from "react-native";
import { Stack } from "expo-router";

import ItemCheckInOutButtons from "../../../components/item/ItemCheckInOutButtons";
import ItemPictures from "../../../components/item/ItemPictures";
import { Text } from "react-native-paper";


const ItemPage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "Item" }} />
            <View style={{ flex: 1 }}>
                <ItemPictures />
                {/** Have to use absolute positioning since the page dots use it :( */}
                <View style={{ paddingLeft: 16, paddingRight: 16, width: "100%" }}>
                    <Text variant="headlineLarge">Item</Text>
                    <Text variant="titleMedium">Item Description</Text>
                    <ItemCheckInOutButtons />
                </View>
            </View>
        </>
    )
}

export default ItemPage;