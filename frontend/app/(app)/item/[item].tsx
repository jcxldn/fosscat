// Based on https://github.com/callstack/react-native-pager-view/blob/master/example/src/PaginationDotsExample.tsx

import React from "react";
import { View } from "react-native";
import { Stack } from "expo-router";

import ItemPictures from "../../../components/item/ItemPictures";


const ItemPage = () => {
    return (
        <>
            <Stack.Screen options={{ title: "Item" }} />
            <View style={{ flex: 1 }}>
                <ItemPictures />
            </View>
        </>
    )
}

export default ItemPage;