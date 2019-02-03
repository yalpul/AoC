#include <iostream>
#include <set>

static const unsigned short limit = 16;

unsigned long int single_word(unsigned char nums[limit])
{
  unsigned long int word = 0lu;
  for (int i = 0; i < limit; i++)
    word |= static_cast<unsigned long int>(nums[i]) << i * 4;
  return word;
}

void distribute(unsigned char *nums)
{
  int idx = 0;
  for (int i = 1; i < limit; i++)
    if (nums[i] > nums[idx])
      idx = i;

  int num = nums[idx];
  nums[idx++] = 0;
  while (num-- > 0)
    nums[idx++ % limit]++;
}

int main()
{
  unsigned char nums[limit];
  unsigned long int word, count = 0;
  std::set<unsigned long int> history;

  for (unsigned short i = 0, n; i < limit; i++)
    std::cin >> n, nums[i] = n;

  word = single_word(nums);
  while (history.find(word) == history.end())
  {
    history.insert(word);
    distribute(nums);
    word = single_word(nums);
    count++;
  }
  std::cout << count << std::endl;
}
