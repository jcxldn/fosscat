import React from "react";
import { useColorScheme } from 'react-native';

import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client';
import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { Slot, Stack } from 'expo-router';
import { PaperProvider } from "react-native-paper";
import { AuthSessionProvider } from "../components/AuthenticationContext";




const RootLayout = () => {
    // https://www.apollographql.com/docs/react/integrations/react-native/
    const gqlClient = new ApolloClient({

        uri: "http://localhost:8080/graphql",

        cache: new InMemoryCache()

    });

    return <RootLayoutNav gqlClient={gqlClient} />
}

const RootLayoutNav = ({ gqlClient }: { gqlClient: ApolloClient<any> }) => {

    const colorScheme = useColorScheme();

    return (
        <ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
            <PaperProvider>
                <ApolloProvider client={gqlClient}>
                    <AuthSessionProvider>
                        <Slot />
                    </AuthSessionProvider>
                </ApolloProvider>
            </PaperProvider>
        </ThemeProvider>
    )
}

export default RootLayout;