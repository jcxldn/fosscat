import { gql } from "@apollo/client";
import { LoginResponse } from "../mutations/login";

interface RefreshTokenData {
    refreshToken: LoginResponse
}

const REFRESH_TOKEN = gql`
    query RefreshToken {
        refreshToken {
            success
            jwt
        }
    }
`
export {
    REFRESH_TOKEN, RefreshTokenData
}