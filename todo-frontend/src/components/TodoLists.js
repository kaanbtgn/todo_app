import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import {
  Container,
  Paper,
  List,
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
  IconButton,
  Button,
  TextField,
  Typography,
  Box,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Checkbox,
  AppBar,
  Toolbar,
} from '@mui/material';
import {
  Delete as DeleteIcon,
  Add as AddIcon,
  Edit as EditIcon,
  ExitToApp as LogoutIcon,
} from '@mui/icons-material';
import { useAuth } from '../context/AuthContext';

function TodoLists() {
  const [lists, setLists] = useState([]);
  const [selectedList, setSelectedList] = useState(null);
  const [newListName, setNewListName] = useState('');
  const [newItemContent, setNewItemContent] = useState('');
  const [isAddListDialogOpen, setIsAddListDialogOpen] = useState(false);
  const [isAddItemDialogOpen, setIsAddItemDialogOpen] = useState(false);
  const { logout } = useAuth();
  const navigate = useNavigate();

  const token = localStorage.getItem('token');
  const axiosInstance = axios.create({
    baseURL: 'http://localhost:8080',
    headers: { Authorization: `Bearer ${token}` },
  });

  useEffect(() => {
    fetchLists();
  }, []);

  const fetchLists = async () => {
    try {
      const response = await axiosInstance.get('/api/todos');
      setLists(response.data || []);
    } catch (error) {
      console.error('Error fetching lists:', error);
      setLists([]);
    }
  };

  const handleAddList = async () => {
    try {
      await axiosInstance.post('/api/todos', { name: newListName });
      setNewListName('');
      setIsAddListDialogOpen(false);
      fetchLists();
    } catch (error) {
      console.error('Error adding list:', error);
    }
  };

  const handleDeleteList = async (listId) => {
    try {
      await axiosInstance.delete(`/api/todos/${listId}`);
      fetchLists();
      if (selectedList?.id === listId) {
        setSelectedList(null);
      }
    } catch (error) {
      console.error('Error deleting list:', error);
    }
  };

  const handleAddItem = async () => {
    try {
      console.log('Adding item to list:', selectedList.id);
      console.log('Item content:', newItemContent);

      const response = await axiosInstance.post(`/api/todos/${selectedList.id}/items`, {
        content: newItemContent,
        is_completed: false,
      });
      console.log('Add item response:', response.data);

      setNewItemContent('');
      setIsAddItemDialogOpen(false);
      
      // Fetch the updated list with items
      const listResponse = await axiosInstance.get(`/api/todos/${selectedList.id}`);
      console.log('Updated list response:', listResponse.data);
      setSelectedList({
        ...listResponse.data,
        items: listResponse.data.items || [],
      });
      
      // Update the lists to reflect the changes
      const listsResponse = await axiosInstance.get('/api/todos');
      setLists(listsResponse.data || []);
    } catch (error) {
      console.error('Error adding item:', error);
      if (error.response) {
        console.error('Error response:', error.response.data);
      }
    }
  };

  const handleToggleItem = async (itemId, isCompleted) => {
    try {
      await axiosInstance.put(`/api/todos/${selectedList.id}/items/${itemId}`, {
        is_completed: !isCompleted,
      });
      const response = await axiosInstance.get(`/api/todos/${selectedList.id}`);
      setSelectedList(response.data);
      fetchLists();
    } catch (error) {
      console.error('Error toggling item:', error);
    }
  };

  const handleDeleteItem = async (itemId) => {
    try {
      await axiosInstance.delete(`/api/todos/${selectedList.id}/items/${itemId}`);
      const response = await axiosInstance.get(`/api/todos/${selectedList.id}`);
      setSelectedList(response.data);
      fetchLists();
    } catch (error) {
      console.error('Error deleting item:', error);
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  const handleListClick = async (list) => {
    try {
      const response = await axiosInstance.get(`/api/todos/${list.id}`);
      setSelectedList({
        ...response.data,
        items: response.data.items || [],
      });
    } catch (error) {
      console.error('Error fetching list details:', error);
    }
  };

  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Todo Lists
          </Typography>
          <IconButton color="inherit" onClick={handleLogout}>
            <LogoutIcon />
          </IconButton>
        </Toolbar>
      </AppBar>
      <Container sx={{ mt: 4 }}>
        <Box display="flex" gap={2}>
          <Paper sx={{ width: '40%', p: 2 }}>
            <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
              <Typography variant="h6">Lists</Typography>
              <Button
                startIcon={<AddIcon />}
                onClick={() => setIsAddListDialogOpen(true)}
                variant="contained"
                size="small"
              >
                Add List
              </Button>
            </Box>
            <List>
              {lists.map((list) => (
                <ListItem
                  key={list.id}
                  component="button"
                  selected={selectedList?.id === list.id}
                  onClick={() => handleListClick(list)}
                >
                  <ListItemText
                    primary={list.name}
                    secondary={`Completion: ${list.completion_percentage}%`}
                  />
                  <ListItemSecondaryAction>
                    <IconButton
                      edge="end"
                      aria-label="delete"
                      onClick={(e) => {
                        e.stopPropagation();
                        handleDeleteList(list.id);
                      }}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>
              ))}
            </List>
          </Paper>

          {selectedList && (
            <Paper sx={{ width: '60%', p: 2 }}>
              <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
                <Typography variant="h6">{selectedList.name} - Items</Typography>
                <Button
                  startIcon={<AddIcon />}
                  onClick={() => setIsAddItemDialogOpen(true)}
                  variant="contained"
                  size="small"
                >
                  Add Item
                </Button>
              </Box>
              <List>
                {(selectedList.items || []).map((item) => (
                  <ListItem key={item.id}>
                    <Checkbox
                      checked={item.is_completed}
                      onChange={() => handleToggleItem(item.id, item.is_completed)}
                    />
                    <ListItemText primary={item.content} />
                    <ListItemSecondaryAction>
                      <IconButton
                        edge="end"
                        aria-label="delete"
                        onClick={() => handleDeleteItem(item.id)}
                      >
                        <DeleteIcon />
                      </IconButton>
                    </ListItemSecondaryAction>
                  </ListItem>
                ))}
              </List>
            </Paper>
          )}
        </Box>

        <Dialog open={isAddListDialogOpen} onClose={() => setIsAddListDialogOpen(false)}>
          <DialogTitle>Add New List</DialogTitle>
          <DialogContent>
            <TextField
              autoFocus
              margin="dense"
              label="List Name"
              fullWidth
              value={newListName}
              onChange={(e) => setNewListName(e.target.value)}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setIsAddListDialogOpen(false)}>Cancel</Button>
            <Button onClick={handleAddList} variant="contained">
              Add
            </Button>
          </DialogActions>
        </Dialog>

        <Dialog open={isAddItemDialogOpen} onClose={() => setIsAddItemDialogOpen(false)}>
          <DialogTitle>Add New Item</DialogTitle>
          <DialogContent>
            <TextField
              autoFocus
              margin="dense"
              label="Item Content"
              fullWidth
              value={newItemContent}
              onChange={(e) => setNewItemContent(e.target.value)}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setIsAddItemDialogOpen(false)}>Cancel</Button>
            <Button onClick={handleAddItem} variant="contained">
              Add
            </Button>
          </DialogActions>
        </Dialog>
      </Container>
    </>
  );
}

export default TodoLists; 