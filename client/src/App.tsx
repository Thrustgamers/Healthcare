import type React from 'react'
import { server } from "../utils/server"
import {useAsync} from 'react-use';
import Login from './login/login';
import Loading from './loading/loading';
import Dashboard from './dashboard/dashboard';


const App: React.FC = () => {

  const ServerConnection = new server()

  const loggedin = useAsync(async () => {
    const result = await ServerConnection.auth.checkStatus()
    return result
  }, []);

  console.log(loggedin)

  


  return (
    <>
      {loggedin.loading ? (
        < Loading color="dark" />
      ) : loggedin.value ? (
        < Dashboard />
      ) : (
        <Login />
      )}
    
    </>
  )
}

export default App