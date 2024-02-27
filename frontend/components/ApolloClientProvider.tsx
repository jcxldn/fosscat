import React from "react";

import { ApolloClient, InMemoryCache, ApolloProvider, gql, NormalizedCacheObject } from '@apollo/client';

const ApolloClientContext = React.createContext<{
    setServerUri: (uri: string) => void,
    createClient: () => void,
    serverUri: string,
    client: ApolloClient<NormalizedCacheObject> | null
}>({
    // Default values
    setServerUri: (uri) => null,
    createClient: () => null,
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

export function ApolloClientProvider(props: React.PropsWithChildren) {
    const [serverUriState, setServerUriState] = React.useState("")
    const [clientState, setClientState] = React.useState(new ApolloClient({ cache: new InMemoryCache() }))

    return (
        <ApolloClientContext.Provider value={{
            setServerUri: (uri) => {
                setServerUriState(uri)
            },
            createClient: () => {
                // https://www.apollographql.com/docs/react/integrations/react-native/
                setClientState(new ApolloClient({

                    uri: `${serverUriState}/graphql`,

                    cache: new InMemoryCache()

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