#include <iostream>
#include <fstream>
#include <vector>

int instruction(std::vector<int> &vec)
{
  int idx = 0, size = vec.size(), count = 0;
  while (idx < size)
  {
    int old = idx;
    idx += vec[idx];
    if (vec[old] >= 3)
      vec[old]--;
    else
      vec[old]++;
    count++;
  }
  return count;
}

int main(int argc, char **argv)
{
  std::ifstream infile(argv[1]);
  std::vector<int> vec;

  while (infile.good())
  {
    int num=0;
    infile >> num;
    if (infile.eof())
      break;
    vec.push_back(num);
  }
  int count = instruction(vec);
  std::cout << count << std::endl;
}
