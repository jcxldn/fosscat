import { gql } from "@apollo/client";
import { LoginResponse } from "./login";

interface LoginRefreshData {
    loginRefresh: LoginResponse
}

const LOGIN_REFRESH = gql`
    mutation LoginRefresh($refresh: String!) {
        loginRefresh(refresh: $refresh) {
            success
            jwt
        }
    }
`
export { LOGIN_REFRESH, LoginRefreshData };