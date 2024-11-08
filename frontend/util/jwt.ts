import { jwtDecode } from "jwt-decode";

const isJwtExpired = (jwtStr: (string | null)): boolean => {
    if (jwtStr) {
        try {
            // Parse token
            const decoded = jwtDecode(jwtStr);

            // Parse expiry time, falling back to epoch 0
            // x1000 to convert to ms (epoch)
            const expiry = new Date(0)
            expiry.setUTCSeconds(decoded.exp || 0)

            return Date.now() > expiry.getTime()
        } catch (error) {
            // Error decoding, display the error and return true for expired.
            console.error(error)
            return true
        }
    } else {
        console.warn("[isJwtExpired] jwt not specified.")
        // no token specified, return true for expired.
        return true
    }
}

export {
    isJwtExpired
}