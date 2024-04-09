// Add better error handlers when using expo-dev-client
// https://docs.expo.dev/bare/install-dev-builds-in-bare/#add-better-error-handlers
import 'expo-dev-client';

import React from "react";
import { useColorScheme } from 'react-native';

import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { Slot, Stack } from 'expo-router';
import { PaperProvider } from "react-native-paper";
import { AuthSessionProvider } from "../components/AuthenticationContext";
import { ApolloClientProvider } from "../components/ApolloClientProvider";

// Polyfill base-64 for jwt-decode until Hermes 0.74 releases
// https://github.com/auth0/jwt-decode/issues/241#issuecomment-1955965836
import { decode } from 'base-64';
global.atob = decode;



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