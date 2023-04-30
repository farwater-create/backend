import jwt from "jsonwebtoken";
import { environment } from "../lib/environment";
import { API_PERMISSIONS } from "../permissions";
const {JWT_SIGNING_KEY} = environment;
console.log(jwt.sign({
    permissions: API_PERMISSIONS.SERVER_AUTHORIZATION_GRANT | API_PERMISSIONS.SERVER_AUTHORIZATION_VIEW
}, JWT_SIGNING_KEY));