import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntDoctor } from '../../api/models/EntDoctor';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import {Grid,} from '@material-ui/core';
import {
  Header,
  Page,
  pageTheme,
  Link,
 } from '@backstage/core';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
const HeaderCustom = {
  minHeight: '50px',
};
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [doctors, setDoctor] = React.useState<EntDoctor[]>(Array);
 useEffect(() => {
   const getPatients = async () => {
     const res = await api.listDoctor({ limit: 20, offset: 0 });
     setLoading(false);
     setDoctor(res);
   };
   getPatients();
 }, [loading]);
 return (
  <Page theme={pageTheme.service}>
  <Header style={HeaderCustom}
    title={`Patient'Register`}
    subtitle="Some quick intro and links."
  >
     <Grid item xs={12}></Grid>
     <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
  </Header>
  <Grid item xs={12}></Grid>
    <TableContainer component={Paper}>
    <Table className={classes.table} aria-label="simple table">
      <TableHead>
        <TableRow>
          <TableCell align="center">nametitle</TableCell>
          <TableCell align="center">name</TableCell>
          <TableCell align="center">tel</TableCell>
          <TableCell align="center">email</TableCell>
          <TableCell align="center">degree</TableCell>
          <TableCell align="center">department</TableCell>
          <TableCell align="center"><Link component={RouterLink} to="/welcome">
        <Button variant="contained" color="primary">
          Register
        </Button>
      </Link></TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {doctors.map((item:any) => (
          <TableRow key={item.id}>
            
            <TableCell align="center">{item.edges?.nametitle?.nameTitle}</TableCell>
            <TableCell align="center">{item.name}</TableCell>
            <TableCell align="center">{item.tel}</TableCell>
            <TableCell align="center">{item.email}</TableCell>
            <TableCell align="center">{item.edges?.degree?.degree}</TableCell>
            <TableCell align="center">{item.edges?.department?.department}</TableCell>
            <TableCell align="center">
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </TableContainer>
  </Page>
);
		}