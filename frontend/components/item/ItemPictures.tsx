import React from "react";

import { View, Animated, Dimensions, ImageSourcePropType } from "react-native";
import PagerView, { PagerViewOnPageScrollEventData } from 'react-native-pager-view';
import { Card } from "react-native-paper";
import { ScalingDot } from 'react-native-animated-pagination-dots';

// Create an Animated version of the PagerView
const AnimatedPagerView = Animated.createAnimatedComponent(PagerView);

type PhotoEntry = {
    key: number,
    source: ImageSourcePropType
}

// Test data before we integrate it into the production DB
const testData: PhotoEntry[] = [
    {
        key: 1,
        source: { uri: "https://picsum.photos/1024" }
    },
    {
        key: 2,
        source: { uri: "https://picsum.photos/1024" }
    }
]

const PhotoCard = ({ entry }: { entry: PhotoEntry }) => {
    return (
        <View
            key={entry.key}
            style={{
                padding: 16
            }}
        >
            <Card>
                <Card.Cover source={entry.source} />
            </Card>
        </View>
    )
}

const ItemPictures = () => {
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
        <View style={{ flex: 1 }}>
            <AnimatedPagerView
                initialPage={0}
                style={{ flex: 1 }}
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

            />
        </View>
    )

}

export default ItemPictures;