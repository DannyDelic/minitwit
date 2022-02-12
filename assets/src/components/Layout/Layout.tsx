import { Fragment } from 'react';
import NavigationBar from "./NavigationBar";

const Layout = (props: { children: React.ReactNode}) => {
  return (
    <Fragment>
        <NavigationBar/>
        <main>
            {props.children}
        </main>
    </Fragment>
  );
};

export default Layout;
