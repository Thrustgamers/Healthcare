import axios from 'axios'
import { useAtom } from 'jotai'
import { SessionData } from '../atoms'
import { config } from 'dotenv'
config()

const key = process.env.SECRET_KEY
const [value] = useAtom(SessionData)

export const request = axios.create({
    baseURL: 'http://127.0.0.1:3000',
    timeout: 1000,
    withCredentials: true,
    headers: {'code ': key },
    proxy: {
        protocol: 'http',
        host: '127.0.0.1',
        port: 3000,
    },
})