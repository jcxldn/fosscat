import React from "react";

import { useStorageState } from '../util/useStorageState';
import { useMutation } from "@apollo/client";
import { LOGIN, LoginData } from "../gql/mutations/login";
import { useApolloClient } from "./ApolloClientProvider";
import { Text } from "react-native";

const AuthContext = React.createContext<{
    login: (email: string, password: string) => Promise<LoginData | null | undefined>,
    logout: () => void,
    jwt: string | null,
    isLoading: boolean,
}>({
    login: async (email: string, password: string) => null,
    logout: () => null,
    jwt: null,
    isLoading: false
})

// Define a hook to be called within SessionProvider to access the session
export function useSession() {
    const value = React.useContext(AuthContext)

    if (process.env.NODE_ENV !== 'production') {
        if (!value) {
            throw new Error('useSession must be wrapped in a <SessionProvider />');
        }
    }

    return value;
}

export function AuthSessionProvider(props: React.PropsWithChildren) {
    const [[isLoading, jwt], setSession] = useStorageState("session")

    const { client } = useApolloClient();

    if (client) {
        const [login, { data, loading, error }] = useMutation<LoginData>(LOGIN);
        return (
            <AuthContext.Provider value={{
                login: async (email: string, password: string) => {
                    await login({ variables: { email, password } })

                    if (loading) {
                        console.error("Still loading after await fulfilled")
                    }

                    if (error) {
                        console.error(error)
                    }

                    if (data?.login.jwt) {
                        setSession(data.login.jwt)
                        console.log("JWT saved in session")
                    }
                    return data
                },
                logout: () => {
                    console.log("logout called")
                },
                jwt,
                isLoading
            }}>
                {props.children}
            </AuthContext.Provider>
        )
    } else {
        // TODO: make a prettier error message
        console.error("Apollo client not available, was it initialized?")
        return <Text>Apollo client not available</Text>
    }
}