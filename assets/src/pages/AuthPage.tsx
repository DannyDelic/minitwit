import AuthForm from '../components/Auth/AuthForm';

interface AuthPageProps {
  fromSignUp?: boolean;
}

const AuthPage = (props: AuthPageProps) => {
  const {fromSignUp} = props;
  return <AuthForm fromSignUp={fromSignUp}/>;
};

export default AuthPage;
