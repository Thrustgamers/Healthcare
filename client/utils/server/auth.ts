import { useAtom } from 'jotai';
import { SessionData } from '../atoms';
import { request } from './request';

type loginData = {
    employeeId: string
    password: string
}

type serverData = {

}

export class auth {

    login(data: loginData) {
        return this.loginHandler(data)
    }

    logout() {
        this.logoutHandler()
    }

    checkStatus() {
        this.checkStatusHandler()
    }

    private async loginHandler (data: loginData) {

        request.post('/login', data).finally((response) => {

        })



    }

    private async logoutHandler () {
        const [value] = useAtom(SessionData)

        const status = await request.post('/logout', value)

        return ""

    }

    private async checkStatusHandler() {

        const [value] = useAtom(SessionData)

        const status = await request.post('/status', value)

        return ""
    }

}