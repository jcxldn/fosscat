import React from "react";
import { useColorScheme } from 'react-native';

import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { Slot, Stack } from 'expo-router';
import { PaperProvider } from "react-native-paper";
import { AuthSessionProvider } from "../components/AuthenticationContext";
import { ApolloClientProvider } from "../components/ApolloClientProvider";




const RootLayout = () => {
    return <RootLayoutNav />
}

const RootLayoutNav = () => {

    const colorScheme = useColorScheme();

    return (
        <ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
            <PaperProvider>
                <ApolloClientProvider>
                    <AuthSessionProvider>
                        <Slot />
                    </AuthSessionProvider>
                </ApolloClientProvider>
            </PaperProvider>
        </ThemeProvider>
    )
}

export default RootLayout;