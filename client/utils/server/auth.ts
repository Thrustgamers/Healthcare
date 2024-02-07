import { useAtom } from 'jotai';
import { SessionData } from '../atoms';
// import { request } from './request';

type loginData = {
    employeeId: string
    password: string
}

export class auth {

    login(data: loginData) {
        return this.loginhandler(data)
    }

    logout() {
        this.logouthandler()
    }

    private async loginhandler (data: loginData) {



    }

    private async logouthandler () {
        const [value] = useAtom(SessionData)

        // const status = await request.get()

        return ""

    }

}