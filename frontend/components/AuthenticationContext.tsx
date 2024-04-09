import React from "react";

import { useStorageState } from '../util/useStorageState';
import { useLazyQuery, useMutation, useQuery } from "@apollo/client";
import { LOGIN, LoginData, LoginResponse } from "../gql/mutations/login";
import { useApolloClient } from "./ApolloClientProvider";
import { Text } from "react-native";
import { isJwtExpired } from "../util/jwt";
import { REFRESH_TOKEN } from "../gql/queries/refreshToken";
import { LOGIN_REFRESH, LoginRefreshData } from "../gql/mutations/loginRefresh";


type JwtToken = {
    jwt: string | null,
    isLoading: boolean
} | null

const AuthContext = React.createContext<{
    login: (email: string, password: string) => Promise<LoginData | null | undefined>,
    logout: () => void,
    getCreds: () => Promise<boolean>,
    access: JwtToken,
    refresh: JwtToken,
}>({
    login: async (email: string, password: string) => null,
    logout: () => null,
    getCreds: async () => false,
    access: null,
    refresh: null
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

    const [[isAccessJwtLoading, accessJwt], setAccess] = useStorageState("access")
    const [[isRefreshJwtLoading, refreshJwt], setRefresh] = useStorageState("refresh")

    const { client, createClient } = useApolloClient();

    const [login, { data, loading, error }] = useMutation<LoginData>(LOGIN);
    const [getRefreshToken, refr] = useLazyQuery(REFRESH_TOKEN);

    // React state is not available on the first render, so render a placeholder on the first render
    // On subsequent renders, react state will have finished loading.
    // Required for login processes to work, otherwise jwts are blank -> isexpiredjwt fails -> always goes to login screen

    // Return lines must be after all hooks.
    if (isAccessJwtLoading || isRefreshJwtLoading) {
        return <Text>Waiting for tokens to be fetched...</Text>
    }

    if (client) {

        const _refreshToken = async () => {

            await getRefreshToken();

            const res = refr.data.refreshToken as LoginResponse

            if (error) {
                console.error(error)
            }

            if (res.success) {
                setRefresh(res.jwt)
                console.log("[AuthContext] Got new refreshToken")
            }
        }

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

                    if (data?.login.success == true) {
                        setAccess(data.login.jwt)
                        createClient(accessJwt!)
                        console.log("[AuthContext/login] Got new accessToken")
                        // Get and save a refresh token 
                        await _refreshToken()
                    }
                    return data
                },
                logout: () => {
                    console.log("logout called")
                },
                getCreds: async () => {
                    // Check if the access token is not present or has expired.
                    if (isJwtExpired(accessJwt)) {
                        console.log("[AuthContext/getCreds] Access JWT expired")
                        // Condition met. Use the refresh token to get a new access token.

                        // Check that the refresh token has not expired
                        if (!isJwtExpired(refreshJwt)) {
                            console.log("[AuthContext/getCreds] Refresh JWT valid.")
                            // Refresh token is valid. Let's call loginRefresh.

                            // Setup the mutation.
                            const [loginRefresh, { data, loading, error }] = useMutation<LoginRefreshData>(LOGIN_REFRESH);

                            // Call the mutation
                            await loginRefresh({ variables: { refresh: refreshJwt } });

                            if (loading) {
                                console.error("Still loading after await fulfilled")
                                return false // not able to get creds
                            }

                            if (error) {
                                console.error(error)
                                return false // not able to get creds
                            }

                            if (data?.loginRefresh.success) {
                                setAccess(data.loginRefresh.jwt)
                                createClient(accessJwt!)
                                console.log("[AuthContext/getCreds] Got new accessToken")
                                await _refreshToken()
                                return true // able to get creds
                            }
                        }
                    }
                    return true // credentials already valid.

                },
                access: {
                    jwt: accessJwt,
                    isLoading: isAccessJwtLoading
                },
                refresh: {
                    jwt: refreshJwt,
                    isLoading: isRefreshJwtLoading
                }
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