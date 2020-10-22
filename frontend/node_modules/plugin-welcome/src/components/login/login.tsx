import React, { FC ,useEffect} from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import LocalHospitalIcon from '@material-ui/icons/LocalHospital';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDoctor } from '../../api/models'; // import interface User
import Swal from 'sweetalert2'; // alert
const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const Login: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const getDoctor = async () => {
    const res = await api.listDoctor({ limit: 10, offset: 0});
    setDoctors(res);
    console.log(res)
  }
  const EmailhandleChange = (event: any) => {
    setEmail(event.target.value as string);
  }
  const PasswordhandleChange = (event: any) => {
    setPassword(event.target.value as string);
  }
  const [email, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);
  useEffect(() => {
    getDoctor();
  }, []);
  var status = false
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  function login(){
    console.log(email,password)
    doctors.map(item =>{
      if(status == false){
        if((item.email == email) && (item.password == password)){
          console.log('pass');
          window.location.href = "http://localhost:3000/home";
          status = true
        }
        else{
          console.log('fail')
          Toast.fire({
            icon: 'error',
            title: 'username หรือ password ไม่ถูกต้อง',
          });
        }
      }
    })
    status = false
  };
  
  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
        <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
      	<LocalHospitalIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            value={email}
            onChange={EmailhandleChange}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            value={password}
            onChange={PasswordhandleChange}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            onClick={login }
            component={RouterLink} to="/"
          >
            Sign In
          </Button>
      </div>
    </Container>
  );
};
export default Login;