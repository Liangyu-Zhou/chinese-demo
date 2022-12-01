import React from 'react';
import {
    AppBar, Toolbar, IconButton, Typography, Button, makeStyles,
} from '@material-ui/core';
import { GitHub } from '@material-ui/icons';
import MenuIcon from '@material-ui/icons/Menu';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
    icon: {
        marginLeft: theme.spacing(1),
    },
    frontColor:{
       
    }
}));

const Header = () => {
    const classes = useStyles();
    return (
        <AppBar position="sticky" >
            <Toolbar>
                <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
                    <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                    Chinese.org Airdrop
                </Typography>
                <Button color="inherit" href="https://github.com/Liangyu-Zhou/registry-demo">
                    View on Github
                    <GitHub className={classes.icon} />
                </Button>
            </Toolbar>
        </AppBar>
    );
}

export default Header;
