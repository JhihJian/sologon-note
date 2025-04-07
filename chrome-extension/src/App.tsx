import React, { useState } from 'react';
import {
  Box,
  Button,
  TextField,
  ToggleButtonGroup,
  ToggleButton,
  Autocomplete,
  FormControlLabel,
  Switch,
  Paper,
  Typography,
  Chip,
  Stack,
} from '@mui/material';
import { styled } from '@mui/material/styles';

const StyledPaper = styled(Paper)(({ theme }) => ({
  width: '400px',
  padding: theme.spacing(2),
}));

const App: React.FC = () => {
  const [recordType, setRecordType] = useState<string>('我干了');
  const [project, setProject] = useState<string | null>(null);
  const [note, setNote] = useState('');
  const [syncGithub, setSyncGithub] = useState(true);
  const [suggestedTags, setSuggestedTags] = useState(['API', '项目', 'OpenAI']);
  const [selectedCategory, setSelectedCategory] = useState('开发笔记');

  const projects = ['项目A', '项目B', '项目C'];
  const categories = ['开发笔记', '灵感记录', '任务提醒', '其他'];

  const handleRecordTypeChange = (
    event: React.MouseEvent<HTMLElement>,
    newType: string,
  ) => {
    if (newType !== null) {
      setRecordType(newType);
    }
  };

  return (
    <StyledPaper elevation={3}>
      <Stack spacing={2}>
        {/* 记录类型选择 */}
        <ToggleButtonGroup
          value={recordType}
          exclusive
          onChange={handleRecordTypeChange}
          fullWidth
        >
          <ToggleButton value="我干了">我干了</ToggleButton>
          <ToggleButton value="我想到了">我想到了</ToggleButton>
          <ToggleButton value="提醒我">提醒我</ToggleButton>
          <ToggleButton value="做记录">做记录</ToggleButton>
        </ToggleButtonGroup>

        {/* 项目选择 */}
        <Autocomplete
          options={projects}
          value={project}
          onChange={(event, newValue) => setProject(newValue)}
          renderInput={(params) => (
            <TextField {...params} label="项目" variant="outlined" />
          )}
        />

        {/* 笔记标题 */}
        <TextField
          label="笔记标题（自动生成）"
          variant="outlined"
          disabled
        />

        {/* 输入区域 */}
        <TextField
          label="输入区：文字/图像/语音"
          multiline
          rows={4}
          value={note}
          onChange={(e) => setNote(e.target.value)}
          variant="outlined"
        />

        {/* 推荐标签 */}
        <Box>
          <Typography variant="body2" color="text.secondary" gutterBottom>
            🔖 推荐标签
          </Typography>
          <Stack direction="row" spacing={1}>
            {suggestedTags.map((tag) => (
              <Chip key={tag} label={`#${tag}`} />
            ))}
            <Chip label="+" variant="outlined" onClick={() => {}} />
          </Stack>
        </Box>

        {/* 分类建议 */}
        <Autocomplete
          options={categories}
          value={selectedCategory}
          onChange={(event, newValue) => setSelectedCategory(newValue || '')}
          renderInput={(params) => (
            <TextField
              {...params}
              label="🗂️ 分类建议"
              variant="outlined"
            />
          )}
        />

        {/* 同步设置和时间戳 */}
        <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <FormControlLabel
            control={
              <Switch
                checked={syncGithub}
                onChange={(e) => setSyncGithub(e.target.checked)}
              />
            }
            label="📤 同步 GitHub"
          />
          <Typography variant="body2" color="text.secondary">
            ⏱ {new Date().toLocaleString()}
          </Typography>
        </Box>

        {/* 按钮组 */}
        <Box sx={{ display: 'flex', justifyContent: 'center', gap: 2 }}>
          <Button variant="contained" color="primary">
            保存记录
          </Button>
          <Button variant="outlined">
            取消
          </Button>
        </Box>
      </Stack>
    </StyledPaper>
  );
};

export default App; 