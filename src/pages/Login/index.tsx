import React, { FC, useEffect, useRef } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import style from './index.module.css'

const Login: FC = () => {
  const { state } = useLocation()
  const navigate = useNavigate()
  const userNameRef = useRef<HTMLInputElement>(null)
  const pwdRef = useRef<HTMLInputElement>(null)
  useEffect(() => {
    if (userNameRef.current && pwdRef.current) {
      userNameRef.current.value = state ? state.username : null
      pwdRef.current.value = state ? state.password : null
    }
  }, [])

  function handleLogin() {
    // 发送登录请求，得到token
    // 如果得到了token以后，判断登录是否成功，如果成功则跳转到主面
    //将token存起来
    navigate('/home')
  }
  return (
    <div className={style['container']}>
      <div className={style['login-wrapper']}>
        <div className={style['header']}>Login</div>
        <div className={style['form-wrapper']}>
          <input type="text" name="username" placeholder="username" className={style['input-item']} ref={userNameRef} />
          <input type="password" name="password" placeholder="password" className={style['input-item']} ref={pwdRef} />
          <div className={style['btn']} onClick={handleLogin}>
            Login
          </div>
        </div>
        <div className={style['msg']}>
          Don't have account?
          <Link to={'/register'}>Sign up</Link>
        </div>
      </div>
    </div>
  )
}
export default Login
