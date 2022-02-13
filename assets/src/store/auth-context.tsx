import React, { useState } from 'react';

interface IAUthContext {
  token: string | null;
  loggedIn: boolean;
  accountId: string | null;
  userName: string | null;
  email: string | null;
  login: (token: string, accountId: string, username: string, email: string) => void;
  logout: () => void;
}

const AuthContext = React.createContext<IAUthContext>({
  token: null,
  loggedIn: false,
  accountId: "",
  userName: "",
  email: "",
  login: (token: string, accountId: string, username: string, email: string) => {},
  logout: () => {},
});

export const AuthContextProvider = (props: { children: React.ReactNode}) => {
  const tokenKey = 'minitwitToken';
  const emailKey = 'minitwitEmail';
  const accountIdKey = 'minitwitAccountId';
  const usernameKey = 'minitwitUsername';

  const [email, setEmail] = useState(localStorage.getItem(emailKey));
  const [accountId, setAccountId] = useState(localStorage.getItem(accountIdKey));
  const [userName, setUserName] = useState(localStorage.getItem(usernameKey));
  const [token, setToken] = useState(localStorage.getItem(tokenKey));

  const loggedIn = !!token;

  const loginHandler = (token: string, accountid: string, username: string, email: string) => {
    setToken(token);
    setEmail(email);
    setAccountId(""+accountid);
    setUserName(username);
    localStorage.setItem(emailKey, email);
    localStorage.setItem(accountIdKey, accountid);
    localStorage.setItem(usernameKey, username);
    localStorage.setItem(tokenKey, token);
  };

  const logoutHandler = () => {
    setToken(null);
    setEmail("");
    setAccountId("");
    setUserName("");
    localStorage.removeItem(emailKey);
    localStorage.removeItem(accountIdKey);
    localStorage.removeItem(usernameKey);
    localStorage.removeItem(tokenKey);

  };

  const contextValue = {
    token: token,
    loggedIn: loggedIn,
    email: email,
    userName: userName,
    accountId: accountId,
    login: loginHandler,
    logout: logoutHandler,
  };

  return (
    <AuthContext.Provider value={contextValue}>
      {props.children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
