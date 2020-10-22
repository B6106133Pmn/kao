import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme,ContentHeader } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2'; // alert
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
  
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDegree } from '../../api/models/EntDegree'; // import interface Degree
import { EntDepartment } from '../../api/models/EntDepartment'; // import interface Department
import { EntNametitle } from '../../api/models/EntNametitle'; // import interface Nametitle

// header css
const HeaderCustom = {
  minHeight: '175px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

const WelcomePage: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [degrees, setDegree] = React.useState<EntDegree[]>([]);
  const [departments, setDepartment] = React.useState<EntDepartment[]>([]);
  const [nametitles, setNameTitle] = React.useState<EntNametitle[]>([]);
 
// alert setting
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

  const getDegree = async () => {
    const res = await http.listDegree({ limit: 4, offset: 0 });
    setDegree(res);
  };

  const getDepartment = async () => {
    const res = await http.listDepartment({ limit: 7, offset: 0 });
    setDepartment(res);
  };

  const getNameTitle = async () => {
    const res = await http.listNametitle({ limit: 2, offset: 0 });
    setNameTitle(res);
  }
  
  const NamehandleChange = (event: any) => {
    setName(event.target.value as string);
  };
  const EmailhandleChange = (event: any) => {
    setEmail(event.target.value as string);
  };
  const PasswordhandleChange = (event: any) => {
    setPassword(event.target.value as string);
  };
  const TelhandleChange = (event: any) => {
    setTel(event.target.value as string);
  };
  const DegreehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDegreeID(event.target.value as number);
  };
  const NametitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNameTitleID(event.target.value as number);
  };
  const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDepartmentID(event.target.value as number);
  };

  const [name, setName] = React.useState(String);
  const [email, setEmail] = React.useState(String);
  const [password, setPassword] = React.useState(String);
  const [tel, setTel] = React.useState(String);
  const [nametitleID, setNameTitleID] = React.useState(Number);
  const [degreeID, setDegreeID] = React.useState(Number);
  const [departmentID, setDepartmentID] = React.useState(Number);
  
  // Lifecycle Hooks
  useEffect(() => {
    getDegree();
    getDepartment();
    getNameTitle();
    }, []);


  // function save data
  function save() {
    const doctors = {
      name: name,
      email: email,
      password: password,
      tel: tel,
      department: departmentID,
      degree: degreeID,
      nameTitle: nametitleID,
    }
    

    const apiUrl = 'http://localhost:8080/api/v1/doctors';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(doctors),
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === false) {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        }
      });
  }
  
  return (
    
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Doctor register system`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
      <Link component={RouterLink} to ="/">
         Logout
        </Link>
      </Header>
      <Content>
      <ContentHeader title="ระบบลงทะเบียนแพทย์">
        </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={5}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Name title</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="Nametitle"
                  value={nametitleID} // (undefined || '') = ''
                  
                  onChange={NametitlehandleChange}
                >
                  {nametitles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nameTitle}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Name</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="name" 
                  label="" 
                  value = {name}
                  variant="outlined" 
                  onChange={NamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Tel.</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="tel" 
                  label="" 
                  value = {tel}
                  variant="outlined" 
                  onChange={TelhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Email</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="name" 
                  label="" 
                  value = {email}
                  variant="outlined" 
                  onChange={EmailhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Password</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                  id="password" 
                  label="" 
                  value = {password}
                  variant="outlined"
                  type="password"
                  autoComplete="current-password"
                  onChange={PasswordhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                  shrink: true,
                }}
                />
              </FormControl>
            </Grid>   
            <Grid item xs={3}>
              <div className={classes.paper}>Degree</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="degree"
                  value={degreeID } // (undefined || '') = ''
                  onChange={DegreehandleChange}
                >
                  {degrees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.degree}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Department</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="department"
                  value={departmentID} // (undefined || '') = ''
                  onChange={DepartmenthandleChange}
                >
                  {departments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.department}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

           
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
                component={RouterLink}
                        to="/table"
              >
                ลงทะเบียน
              </Button>
            </Grid>
          </Grid>
          
             
        </Container>
        <Link component={RouterLink} to ="/table">
        ลงทะเบียน
        </Link>
      </Content>
    </Page>
  );
};
export default WelcomePage;
