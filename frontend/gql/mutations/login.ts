import { gql } from "@apollo/client";

interface LoginResponse {
    success: boolean,
    jwt: string | null
}

interface LoginData {
    login: LoginResponse
}

const LOGIN = gql`
    mutation Login($email: String!, $password: String!) {
        login(email: $email, password: $password) {
            success
            jwt
        }
    }
`

export {
    LOGIN, LoginData, LoginResponse
};