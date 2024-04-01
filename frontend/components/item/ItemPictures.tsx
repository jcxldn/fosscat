import React from "react";

import { View, Animated, Dimensions, ImageSourcePropType, useColorScheme } from "react-native";
import PagerView, { PagerViewOnPageScrollEventData } from 'react-native-pager-view';
import { Button, Card, Surface } from "react-native-paper";
import { ScalingDot } from 'react-native-animated-pagination-dots';

import { MaterialCommunityIcons } from '@expo/vector-icons';

// Create an Animated version of the PagerView
const AnimatedPagerView = Animated.createAnimatedComponent(PagerView);

type PhotoEntry = {
    key: number,
    source?: ImageSourcePropType,
    component?: React.JSX.Element
}


const PhotoCard = ({ entry }: { entry: PhotoEntry }) => {
    return (
        <View
            key={entry.key}
            style={{
                padding: 16,
            }}
        >
            <Card>
                {entry.source ? <Card.Cover source={entry.source} /> : entry.component ? entry.component : null}
            </Card>
        </View >
    )
}

const ItemPictures = () => {

    const colorScheme = useColorScheme()

    // Test data before we integrate it into the production DB

    // Key is for React only, does not affect order.
    const testData: PhotoEntry[] = [
        {
            key: 1,
            source: { uri: "https://picsum.photos/1024" },
        },
        {
            key: 2,
            source: { uri: "https://picsum.photos/1024" }
        },
        {
            key: -1, component: (
                <>
                    <Button style={{ height: 195 }} onPress={() => alert("add image handler tbd")}>
                        <Surface elevation={0} style={{ height: "100%", alignItems: "center", justifyContent: "center" }}>
                            <MaterialCommunityIcons name="image-plus" size={128} color={colorScheme == "dark" ? "white" : "black"} style={{ opacity: 0.5 }} />
                        </Surface>
                    </Button>
                </>
            )
        }
    ]

    // Variables for AnimatedPagerView
    // From "react-native-pager-view" PaginationDotsExample.tsx
    // <https://github.com/callstack/react-native-pager-view/blob/master/example/src/PaginationDotsExample.tsx>
    const width = Dimensions.get('window').width;
    const ref = React.useRef<PagerView>(null);
    const scrollOffsetAnimatedValue = React.useRef(new Animated.Value(0)).current;
    const positionAnimatedValue = React.useRef(new Animated.Value(0)).current;
    const inputRange = [0, testData.length];
    const scrollX = Animated.add(
        scrollOffsetAnimatedValue,
        positionAnimatedValue
    ).interpolate({
        inputRange,
        outputRange: [0, testData.length * width],
    });

    const onPageScroll = React.useMemo(
        () =>
            Animated.event<PagerViewOnPageScrollEventData>(
                [
                    {
                        nativeEvent: {
                            offset: scrollOffsetAnimatedValue,
                            position: positionAnimatedValue,
                        },
                    },
                ],
                {
                    useNativeDriver: false,
                }
            ),
        // eslint-disable-next-line react-hooks/exhaustive-deps
        []
    );

    return (
        <View style={{ flex: 1, maxHeight: 275 }}>
            <AnimatedPagerView
                initialPage={0}
                style={{ flex: 1, maxHeight: 250 }}
                ref={ref}
                onPageScroll={onPageScroll}
            >
                {/** Set key prop to make React happy */}
                {testData.map(entry => <PhotoCard key={entry.key} entry={entry} />)}
            </AnimatedPagerView>
            <ScalingDot
                data={testData}
                //@ts-ignore
                scrollX={scrollX}

                dotStyle={{
                    width: 7.5,
                    height: 7.5,
                    //backgroundColor: '#347af0',
                    borderRadius: 5,
                    marginHorizontal: 5,
                }}
                containerStyle={{ top: 250 }}
            />
        </View>
    )

}

export default ItemPictures;