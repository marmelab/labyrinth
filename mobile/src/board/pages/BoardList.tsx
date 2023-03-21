import { useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";

import {
  Button,
  List as MuiList,
  ListSubheader,
  ListItem,
  ListItemText,
  Typography,
} from "@mui/material";

import { BoardListItem } from "../BoardTypes";
import { boardRepository } from "../BoardRepository";

import { useUserContext } from "../../user/UserContext";

export function List() {
  const navigate = useNavigate();
  const [user, _] = useUserContext();
  const [boards, setBoards] = useState<BoardListItem[]>([]);

  useEffect(() => {
    boardRepository.list(1).then(setBoards);
  }, [user]);

  return (
    <>
      {user ? (
        <Button onClick={() => navigate("/board/new")} variant="contained">
          Create New Board
        </Button>
      ) : (
        <Typography fontWeight={700}>
          You are not signed in! <br />
          You can spectate these games:
        </Typography>
      )}
      <MuiList
        sx={{ width: "100%", maxWidth: 360, bgcolor: "background.paper" }}
        component="nav"
        aria-labelledby="nested-list-subheader"
        subheader={
          <ListSubheader component="div" id="nested-list-subheader">
            Boards
          </ListSubheader>
        }
      >
        {boards.map((board) => (
          <ListItem key={board.id}>
            <ListItemText
              primary={
                <Link to={`/board/${board.id}/view`}>Board #{board.id}</Link>
              }
            />
          </ListItem>
        ))}
      </MuiList>
    </>
  );
}
