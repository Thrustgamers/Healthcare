import axios from 'axios'
import { config } from 'dotenv'
config()

const key = process.env.SECRET_KEY

const request = axios.create({
    baseURL: 'http://localhost:3000',
    timeout: 1000,
    withCredentials: true,
    headers: {'code ': key}
})


type serverRequest = {
    get(): () => request.get()
}
