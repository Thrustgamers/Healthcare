import { auth } from "./auth";

export class server {
    auth: auth

    constructor() {
        this.auth = new auth()
    }
}