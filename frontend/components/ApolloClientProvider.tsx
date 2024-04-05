import React from "react";

import { ApolloClient, InMemoryCache, ApolloProvider, gql, NormalizedCacheObject } from '@apollo/client';

const ApolloClientContext = React.createContext<{
    setServerUri: (uri: string) => void,
    createClient: (access?: string) => void,
    serverUri: string,
    client: ApolloClient<NormalizedCacheObject> | null
}>({
    // Default values
    setServerUri: (uri) => null,
    createClient: (access) => null,
    serverUri: "",
    client: null
})

// Define a hook to be called within ApolloClientProvider to access the client
export function useApolloClient() {
    const value = React.useContext(ApolloClientContext)

    if (process.env.NODE_ENV !== 'production') {
        if (!value) {
            throw new Error('useApolloClient must be wrapped in a <ApolloClientProvider />');
        }
    }

    return value;
}

function getNoAuthClient(uri: string) {
    return new ApolloClient({
        uri: `${uri}/graphql`,
        cache: new InMemoryCache()
    })
}

export function ApolloClientProvider(props: React.PropsWithChildren) {
    const [serverUriState, setServerUriState] = React.useState("")
    // Use a React state for the ApolloClient, setting the default to a client without auth headers.
    // NB: uri is not set here?
    const [clientState, setClientState] = React.useState(getNoAuthClient(serverUriState))

    return (
        <ApolloClientContext.Provider value={{
            setServerUri: (uri) => {
                setServerUriState(uri)
                setClientState(getNoAuthClient(uri))
            },
            createClient: (access?: string) => {
                // https://www.apollographql.com/docs/react/integrations/react-native/
                setClientState(new ApolloClient({

                    uri: `${serverUriState}/graphql`,

                    cache: new InMemoryCache(),

                    headers: {
                        // Set auth header if access token as passed
                        authorization: access ? `Bearer ${access}` : ""
                    }

                }));
            },
            serverUri: serverUriState,
            client: clientState
        }}>
            {clientState ?
                <ApolloProvider client={clientState}>{props.children}</ApolloProvider> :
                // Client not available, just render children as a fallback
                <>
                    {props.children}
                </>
            }
        </ApolloClientContext.Provider>
    )
}