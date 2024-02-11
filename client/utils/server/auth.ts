import { request } from './request';

type loginData = {
    employeeId: string
    password: string
}

export class auth {

    async login(data: loginData) {
        return await this.loginHandler(data)
    }

    async logout() {
        return await this.logoutHandler()
    }

    async checkStatus() {
        return await this.checkStatusHandler()
    }

    private async loginHandler (data: loginData) {

        request.post('/login', data).then(() => {

        })

    }

    private async logoutHandler () {
        const response = await request.post('/logout');
        return response.status;
    }

    private async checkStatusHandler() {
        const response = await request.post('/statuscheck')

        console.log(response)
        return response.data.data.status === "success"
    }

}