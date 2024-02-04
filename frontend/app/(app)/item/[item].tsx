// Based on https://github.com/callstack/react-native-pager-view/blob/master/example/src/PaginationDotsExample.tsx

import React from "react";
import { View } from "react-native";
import { Stack } from "expo-router";

import ItemPictures from "../../../components/item/ItemPictures";
import { Text } from "react-native-paper";


const ItemPage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "Item" }} />
            <View style={{ flex: 1 }}>
                <ItemPictures />
                <View style={{ flex: 1, padding: 16 }}>
                    <Text variant="headlineLarge">Item</Text>
                    <Text variant="titleMedium">Item Description</Text>
                </View>
            </View>
        </>
    )
}

export default ItemPage;