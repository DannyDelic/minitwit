import React, { useState } from 'react';

interface IAUthContext {
  token: string | null;
  loggedIn: boolean;
  login: (token: string) => void;
  logout: () => void;
}

const AuthContext = React.createContext<IAUthContext>({
  token: null,
  loggedIn: false,
  login: (token: string) => {},
  logout: () => {},
});

export const AuthContextProvider = (props: { children: React.ReactNode}) => {
  const tokenKey = 'minitwitToken';

  const [token, setToken] = useState(localStorage.getItem(tokenKey));

  const loggedIn = !!token;

  const loginHandler = (token: string) => {
    setToken(token);
    localStorage.setItem(tokenKey, token);
  };

  const logoutHandler = () => {
    setToken(null);
    localStorage.removeItem(tokenKey);
  };

  const contextValue = {
    token: token,
    loggedIn: loggedIn,
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
