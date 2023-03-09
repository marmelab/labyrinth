import { useContext, useEffect, type ReactNode } from "react";
import { useNavigate, Link as RouterLink, Outlet } from "react-router-dom";

import {
  AppBar,
  Box,
  Container,
  Link,
  MenuItem,
  Toolbar,
  Typography,
} from "@mui/material";

import { userRepository, UserContext } from "../../user";

const NavLink = function ({
  to,
  onClick,
  children,
}: {
  to?: string;
  onClick?: () => Promise<void>;
  children: ReactNode;
}) {
  const styles = { my: 2, color: "white", display: "block" };
  return (
    <>
      {to && (
        <Link component={RouterLink} to={to} sx={styles}>
          {children}
        </Link>
      )}
      {onClick && (
        <Link onClick={onClick} sx={styles}>
          {children}
        </Link>
      )}
    </>
  );
};

const Layout = () => {
  const [user, setUser] = useContext(UserContext);
  const navigate = useNavigate();

  useEffect(() => {
    userRepository.getIdentity().then(setUser);
  }, []);

  const signOut = async function () {
    await userRepository.signOut();
    setUser(null);
    navigate("/");
  };

  return (
    <>
      <AppBar>
        <Container maxWidth="xl">
          <Toolbar disableGutters>
            <Box flexGrow={1}>
              <MenuItem>
                <NavLink to="/">Home</NavLink>
              </MenuItem>
            </Box>

            <Box>
              <MenuItem>
                {!user ? (
                  <NavLink to="/auth/sign-in">Sign In / Sign Up</NavLink>
                ) : (
                  <>
                    <Typography sx={{ m: 2, color: "white", display: "block" }}>
                      {user.name}
                    </Typography>

                    <NavLink onClick={signOut}>Sign Out</NavLink>
                  </>
                )}
              </MenuItem>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
      <main>
        <Outlet />
      </main>
    </>
  );
};

export default Layout;
